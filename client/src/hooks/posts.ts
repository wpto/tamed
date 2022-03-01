import { useCallback, useState } from 'react'
import { PostPreview } from '../types/render'
import { useFetch } from './fetch'

interface ApiPost {
  id: string
  ctime: string
  tags: string[]
  link: string
}

interface ApiPostList {
  page: number
  pages: number
  total: number
  posts: ApiPost[]
  tags: string[]
}

type JSONResponse = {
  data?: ApiPostList
  errors?: Array<string>
}

export const usePosts = () => {
  const [page, setPage] = useState(0)
  const { request } = useFetch()

  const [loaded, setLoaded] = useState(false)
  const [posts, setPosts] = useState<PostPreview[]>([])
  const [error, setError] = useState<string | null>(null)

  const fetchPosts = useCallback(async (idx: number = 0) => {
    setLoaded(false)
    setError(null)
    const { data, errors }: JSONResponse = await request(
      `/api/posts?offset=${page}`
    )

    if (errors != null) {
      setError(errors.join('\n'))
    }

    if (data != null) {
      setPosts(
        data.posts.map((e) => ({
          thumbUrl: `/media/${e.link}`,
          link: `/post/${e.id}`,
        }))
      )
    }

    setLoaded(true)
    setPage(idx)
  }, [])

  return { posts, loaded, page, fetchPosts, error }
}
