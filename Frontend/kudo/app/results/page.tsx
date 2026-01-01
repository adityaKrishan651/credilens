'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend,
} from 'chart.js'
import { Doughnut, Pie } from 'react-chartjs-2'

ChartJS.register(ArcElement, Tooltip, Legend)

export default function ResultsPage() {
  const router = useRouter()

  const score = 85
  const riskLevel: 'Low' | 'Medium' | 'High' = 'Medium'

  const riskData = {
    low: 60,
    medium: 30,
    high: 10,
  }

//   type RiskLevel = 'Low' | 'Medium' | 'High';

// interface Props {
//   riskLevel: RiskLevel;
// }

// const riskBadgeClass = 
//   riskLevel === 'Low'
//     ? 'bg-success'
//     : riskLevel === 'Medium'
//     ? 'bg-warning'
//     : 'bg-danger';
 
  const riskBadgeClass : Record<string, string> = {
  Low: 'bg-success',
  Medium: 'bg-warning',
  High: 'bg-danger',
};

// Fallback to 'bg-danger' or a default if riskLevel is missing
const currentBg = riskBadgeClass[riskLevel] || 'bg-secondary';
  return (
    <>
      {/* Header */}
      <header className="bg-primary text-white py-3 shadow-sm">
        <div className="container d-flex justify-content-between align-items-center">
          <button
            className="btn btn-outline-light btn-sm"
            onClick={() => router.back()}
          >
            <i className="fas fa-arrow-left me-1" /> Back
          </button>
          <h3 className="mb-0 fw-semibold">Analysis Results</h3>
          <span />
        </div>
      </header>

      {/* Main */}
      <main className="container my-5">
        {/* Score & Risk */}
        <div className="row g-4">
          <div className="col-md-6">
            <div className="card shadow-sm text-center h-100">
              <div className="card-body">
                <h4>Credibility Score</h4>

                <div style={{ maxWidth: 220, margin: '0 auto' }}>
                  <Doughnut
                    data={{
                      datasets: [
                        {
                          data: [score, 100 - score],
                          backgroundColor: ['#198754', '#e9ecef'],
                          borderWidth: 0,
                        },
                      ],
                    }}
                    options={{
                      cutout: '75%',
                      plugins: { legend: { display: false } },
                    }}
                  />
                </div>

                <p className="mt-3 fs-5">
                  <strong>{score}</strong> / 100
                </p>
                <small className="text-muted">
                  Overall trustworthiness assessment
                </small>
              </div>
            </div>
          </div>

          <div className="col-md-6">
            <div className="card shadow-sm h-100">
              <div className="card-body text-center">
                <h4>Risk Level</h4>

                <span
                  className={`badge fs-6 mb-3 d-inline-block ${riskBadgeClass}`}
                >
                  {riskLevel}
                </span>

                <Pie
                  data={{
                    labels: ['Low', 'Medium', 'High'],
                    datasets: [
                      {
                        data: [
                          riskData.low,
                          riskData.medium,
                          riskData.high,
                        ],
                        backgroundColor: [
                          '#198754',
                          '#ffc107',
                          '#dc3545',
                        ],
                      },
                    ],
                  }}
                />
              </div>
            </div>
          </div>
        </div>

        {/* Indicators */}
        <div className="row mt-5">
          <div className="col-12">
            <h4>Fraud Indicators</h4>
            <div
              className="accordion shadow-sm rounded"
              id="indicatorsAccordion"
            >
              <div className="accordion-item">
                <h2 className="accordion-header">
                  <button
                    className="accordion-button"
                    data-bs-toggle="collapse"
                    data-bs-target="#indicator1"
                  >
                    <i className="fas fa-exclamation-triangle text-warning me-2" />
                    Suspicious Links Detected
                  </button>
                </h2>
                <div
                  id="indicator1"
                  className="accordion-collapse collapse show"
                >
                  <div className="accordion-body">
                    Found <strong>3 links</strong> pointing to unverified
                    domains.
                    <a href="#" className="ms-2">
                      View details
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Evidence */}
        <div className="row mt-5">
          <div className="col-12">
            <h4>Supporting Evidence</h4>
            <div className="card shadow-sm">
              <div className="card-body">
                <ul className="list-group list-group-flush">
                  <li className="list-group-item">
                    <strong>Source Check:</strong> No matches in known scam
                    databases
                  </li>
                  <li className="list-group-item">
                    <strong>Content Match:</strong> 95% similarity with
                    verified sources
                  </li>
                  <li className="list-group-item">
                    <strong>Recommendation:</strong> Verify via official
                    communication channels
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="text-center py-4 bg-white border-top">
        <button className="btn btn-primary me-2">
          <i className="fas fa-rotate-right me-1" /> Re-Analyze
        </button>
        <button className="btn btn-outline-secondary">
          <i className="fas fa-share-nodes me-1" /> Share Results
        </button>
      </footer>
    </>
  )
}
