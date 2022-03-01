import React, { useEffect, useState } from 'react'
import { Route, Routes } from 'react-router-dom'
import styles from './App.module.css'
import { PostList } from './components/PostList/PostList'
import { usePosts } from './hooks/posts'
import { Post } from './routes/Post/Post'

export const App: React.FC = () => {
  const { fetchPosts, loaded, posts } = usePosts()
  useEffect(() => {
    fetchPosts()
  }, [])
  return (
    <div>
      <div>
        {!loaded ? <div>loading...</div> : null}
        <PostList posts={posts} />
      </div>
      <Routes>
        <Route path="/post/:postId" element={<Post />} />
      </Routes>
    </div>
  )
}
