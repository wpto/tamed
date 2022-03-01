import React from 'react'
import { useParams } from 'react-router-dom'
import { PostView } from '../../components/PostView/PostView'
import { usePost } from '../../hooks/post'

export const Post: React.FC = () => {
  const params = useParams()
  const { post, loaded } = usePost(params.postId)
  console.log(post)
  return <div>{loaded ? <PostView fullUrl={post.fullUrl} /> : null}</div>
}
