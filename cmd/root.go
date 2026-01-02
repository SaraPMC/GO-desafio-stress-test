package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Ferramenta CLI para testes de carga em serviços web",
	Long: `Stress Test é uma ferramenta CLI em Go para realizar testes de carga
em serviços web. Você pode especificar a URL, número total de requisições
e quantidade de chamadas simultâneas.`,
	RunE: runStressTest,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
	rootCmd.Flags().IntVar(&requests, "requests", 0, "Número total de requests (obrigatório)")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas (padrão: 1)")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
}

type RequestResult struct {
	StatusCode int
	Duration   time.Duration
}

type Report struct {
	TotalRequests   int64
	SuccessRequests int64
	FailedRequests  int64
	TotalTime       time.Duration
	StatusCodeCount map[int]int64
	AverageDuration time.Duration
	MinDuration     time.Duration
	MaxDuration     time.Duration
}

func runStressTest(cmd *cobra.Command, args []string) error {
	fmt.Printf("\n🚀 Iniciando teste de carga\n")
	fmt.Printf("📍 URL: %s\n", url)
	fmt.Printf("📊 Total de requests: %d\n", requests)
	fmt.Printf("⚡ Concorrência: %d\n\n", concurrency)

	startTime := time.Now()
	report := executeStressTest()
	report.TotalTime = time.Since(startTime)

	printReport(report)

	return nil
}

func executeStressTest() *Report {
	report := &Report{
		StatusCodeCount: make(map[int]int64),
		MinDuration:     time.Hour,
	}

	var wg sync.WaitGroup
	requestsChan := make(chan struct{}, concurrency)
	resultsChan := make(chan RequestResult, 100)
	var totalDuration int64

	// Goroutine para coletar resultados
	go func() {
		for result := range resultsChan {
			report.StatusCodeCount[result.StatusCode]++
			if result.StatusCode == 200 {
				atomic.AddInt64(&report.SuccessRequests, 1)
			} else {
				atomic.AddInt64(&report.FailedRequests, 1)
			}

			if result.Duration < report.MinDuration {
				report.MinDuration = result.Duration
			}
			if result.Duration > report.MaxDuration {
				report.MaxDuration = result.Duration
			}

			atomic.AddInt64(&totalDuration, int64(result.Duration))
		}
	}()

	// Iniciar workers
	for i := 0; i < concurrency; i++ {
		go worker(requestsChan, resultsChan, &wg)
	}

	// Enviar requisições
	wg.Add(requests)
	go func() {
		for i := 0; i < requests; i++ {
			requestsChan <- struct{}{}
		}
		close(requestsChan)
	}()

	wg.Wait()
	close(resultsChan)

	report.TotalRequests = int64(requests)
	if requests > 0 {
		report.AverageDuration = time.Duration(totalDuration / int64(requests))
	}

	return report
}

func worker(requestsChan <-chan struct{}, resultsChan chan<- RequestResult, wg *sync.WaitGroup) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	for range requestsChan {
		defer wg.Done()

		startTime := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(startTime)

		if err != nil {
			resultsChan <- RequestResult{
				StatusCode: 0,
				Duration:   duration,
			}
			continue
		}

		resp.Body.Close()

		resultsChan <- RequestResult{
			StatusCode: resp.StatusCode,
			Duration:   duration,
		}
	}
}

func printReport(report *Report) {
	fmt.Println("\n" + "="*60)
	fmt.Println("📋 RELATÓRIO DE TESTE DE CARGA")
	fmt.Println("="*60 + "\n")

	fmt.Printf("⏱️  Tempo total: %v\n", report.TotalTime)
	fmt.Printf("📊 Total de requests: %d\n", report.TotalRequests)
	fmt.Printf("✅ Requests com status 200: %d (%.2f%%)\n",
		report.SuccessRequests,
		float64(report.SuccessRequests)/float64(report.TotalRequests)*100)
	fmt.Printf("❌ Requests com falha: %d\n\n", report.FailedRequests)

	fmt.Println("📈 Estatísticas de Duração:")
	fmt.Printf("   ⚡ Mínimo: %v\n", report.MinDuration)
	fmt.Printf("   ⏱️  Médio: %v\n", report.AverageDuration)
	fmt.Printf("   🐢 Máximo: %v\n\n", report.MaxDuration)

	fmt.Println("🔢 Distribuição de Códigos HTTP:")
	for statusCode := 100; statusCode < 600; statusCode += 100 {
		for code, count := range report.StatusCodeCount {
			if code >= statusCode && code < statusCode+100 && count > 0 {
				fmt.Printf("   HTTP %d: %d requisições\n", code, count)
			}
		}
	}

	fmt.Printf("\n📊 Taxa de requisições por segundo: %.2f req/s\n",
		float64(report.TotalRequests)/report.TotalTime.Seconds())

	fmt.Println("\n" + "="*60 + "\n")
}
