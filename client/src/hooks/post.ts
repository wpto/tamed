import { useEffect, useState } from 'react'
import { PostViewProps } from '../types/render'
import { useFetch } from './fetch'

type JSONResponse = {
  errors?: Array<string>
  data?: {
    id: string
    ctime: string
    tags: string[]
    link: string
  }
}

export const usePost = (postId: string = '') => {
  const { request } = useFetch()
  const [post, setPost] = useState<PostViewProps>({ fullUrl: '' })
  const [loaded, setLoaded] = useState(false)

  useEffect(() => {
    ;(async () => {
      setLoaded(false)
      if (postId.length !== 0) {
        const { data }: JSONResponse = await request(`/api/posts/${postId}`)

        if (data) {
          setPost({ fullUrl: `/media/${data.link}` })
        }

        setLoaded(true)
      }
    })()
  }, [postId])

  return { post, loaded }
}
