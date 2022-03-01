import { useState, useCallback } from 'react'

export const useFetch = () => {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<Error | null>(null)

  const request = useCallback(
    async (
      url: RequestInfo,
      method: string = 'GET',
      body: Object | null = null,
      headers: HeadersInit = {}
    ): Promise<{ data: any; errors: string[] }> => {
      setLoading(true)
      try {
        let bodyInit: BodyInit | null = null
        let incHeaders: { 'content-type'?: string } = {}
        if (body) {
          bodyInit = JSON.stringify(body) as BodyInit
          incHeaders['content-type'] = 'application/json;charset=UTF-8'
        }

        const response = await fetch(url, {
          method,
          body: bodyInit,
          headers: {
            ...incHeaders,
            ...headers,
          },
        })

        const data = await response.json()

        if (!response.ok) {
          throw new Error(data.message || 'Something went wrong')
        }

        setLoading(false)
        return { data, errors: [] }
      } catch (e) {
        setLoading(false)
        return { data: null, errors: [String(e)] }
      }
    },
    []
  )

  const clearError = useCallback(() => setError(null), [])

  return { loading, request, error, clearError }
}
