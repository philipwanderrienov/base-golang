import { useState, useEffect } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import './App.css'

function App() {
  const [congregations, setCongregations] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetch('http://localhost:8080/api/v1/congregations')
      .then(response => response.json())
      .then(data => {
        setCongregations(data.data || [])
        setLoading(false)
      })
      .catch(err => {
        setError(err.message)
        setLoading(false)
      })
  }, [])

  if (loading) return <div className="flex justify-center items-center h-screen text-lg">Loading...</div>
  if (error) return <div className="flex justify-center items-center h-screen text-red-600 text-lg">Error: {error}</div>

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-4xl font-bold text-center mb-12 text-gray-950">Congregation Dashboard</h1>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {congregations && congregations.length > 0 ? (
            congregations.map(cong => (
              <Card key={cong.id}>
                <CardHeader>
                  <CardTitle>{cong.name}</CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="space-y-2">
                    <p className="text-sm text-gray-600">
                      <span className="font-semibold">ID:</span> {cong.id}
                    </p>
                    <p className="text-sm text-gray-600">
                      <span className="font-semibold">Location:</span> {cong.location}
                    </p>
                  </div>
                </CardContent>
              </Card>
            ))
          ) : (
            <div className="col-span-3 text-center text-gray-500">
              No congregations found
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default App
