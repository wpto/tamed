import { useCallback, useState } from 'react'
import { useFetch } from './fetch'

export const useUpload = () => {
  const [status, setStatus] = useState<string | null>(null)
  const upload = useCallback((files: FileList) => {
    var data = new FormData()

    for (let i = 0; i < files.length; i++) {
      const f = files[i]
      data.append('upload[]', f)
    }

    setStatus('uploading...')
    try {
      fetch('/api/posts', {
        method: 'POST',
        body: data,
      })
      setStatus('done!')
    } catch (e) {
      setStatus('error: ' + e)
    }
  }, [])

  return { status, upload }
}
