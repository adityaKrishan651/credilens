'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'

type InputType = 'url' | 'text' | 'file'

export default function HomePage() {
  const [inputType, setInputType] = useState<InputType>('url')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(false)
  const [progress, setProgress] = useState(0)

  const router = useRouter()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError(false)
    setProgress(0)

    const interval = setInterval(() => {
      setProgress((prev) => {
        if (prev >= 100) {
          clearInterval(interval)
          setTimeout(() => {
            if (Math.random() > 0.5) {
              router.push('/results')
            } else {
              setLoading(false)
              setError(true)
            }
          }, 500)
          return 100
        }
        return prev + 10
      })
    }, 200)
  }

  return (
    <>
      {/* Hero */}
      <header className="hero text-white text-center py-5">
        <div className="container position-relative">
          <h1 className="display-4">Detect Spam & Fraud Instantly</h1>
          <p className="lead">
            Upload a URL, paste text, or share an image to check credibility.
          </p>
        </div>
      </header>

      {/* Main */}
      <main className="container my-5">
        <div className="row justify-content-center">
          <div className="col-md-8">
            <div className="card shadow-lg">
              <div className="card-header">
                <h3>Choose Your Input Type</h3>
                <div className="btn-group w-100">
                  {(['url', 'text', 'file'] as InputType[]).map((type) => (
                    <button
                      key={type}
                      type="button"
                      className={`btn btn-outline-primary ${
                        inputType === type ? 'active' : ''
                      }`}
                      onClick={() => setInputType(type)}
                    >
                      {type.toUpperCase()}
                    </button>
                  ))}
                </div>
              </div>

              <div className="card-body">
                <form onSubmit={handleSubmit}>
                  {inputType === 'url' && (
                    <div className="input-group mb-3">
                      <span className="input-group-text">
                        <i className="fas fa-link" />
                      </span>
                      <input
                        type="url"
                        className="form-control"
                        placeholder="https://example.com"
                        required
                      />
                    </div>
                  )}

                  {inputType === 'text' && (
                    <div className="input-group mb-3">
                      <span className="input-group-text">
                        <i className="fas fa-edit" />
                      </span>
                      <textarea
                        className="form-control"
                        rows={4}
                        placeholder="Paste your text here..."
                        required
                      />
                    </div>
                  )}

                  {inputType === 'file' && (
                    <div className="input-group mb-3">
                      <input
                        type="file"
                        className="form-control"
                        accept="image/*,.pdf,.txt"
                        required
                      />
                      <span className="input-group-text">
                        <i className="fas fa-upload" />
                      </span>
                    </div>
                  )}

                  <button
                    className="btn btn-primary w-100"
                    disabled={loading}
                  >
                    Analyze
                  </button>
                </form>

                {loading && (
                  <div className="text-center mt-3">
                    <div className="spinner-border text-primary" />
                    <p className="mt-2">Analyzing your input...</p>
                    <div className="progress">
                      <div
                        className="progress-bar progress-bar-striped progress-bar-animated"
                        style={{ width: `${progress}%` }}
                      />
                    </div>
                  </div>
                )}

                {error && (
                  <div className="alert alert-danger mt-3">
                    <i className="fas fa-exclamation-triangle" /> Invalid input.
                    Please try again.
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-light text-center py-3">
        <p>&copy; 2026 Spam & Fraud Detector. <a href="#">Privacy Policy</a></p>
      </footer>
    </>
  )
}