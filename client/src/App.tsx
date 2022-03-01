import React, { useEffect, useState } from 'react'
import styles from './App.module.css'
import { PostList } from './components/PostList/PostList'
import { usePosts } from './hooks/posts'

export const App: React.FC = () => {
  const { fetchPosts, loaded, posts } = usePosts()
  useEffect(() => {
    fetchPosts()
  }, [])
  return (
    <div>
      {!loaded ? <div>loading...</div> : null}
      <PostList posts={posts} />
    </div>
  )
}
