import React, { useCallback, useEffect, useState } from 'react'
import { Link, Route, Routes } from 'react-router-dom'
import styles from './App.module.css'
import { Dropzone } from './components/Dropzone/Dropzone'
import { PostList } from './components/PostList/PostList'
import { usePosts } from './hooks/posts'
import { Post } from './routes/Post/Post'

export const App: React.FC = () => {
  const { fetchPosts, loaded, posts } = usePosts()
  useEffect(() => {
    fetchPosts()
  }, [])

  const dropUpdate = useCallback(() => {
    fetchPosts()
  }, [])

  return (
    <div className={styles.app}>
      <header className={styles.header}>
        <ul className={styles.nav}>
          <li className={styles.navLink}>
            <Link to="/">Posts</Link>
          </li>
          <li className={styles.navLink}>
            <Link to="/">Posts</Link>
          </li>
        </ul>
      </header>
      <div className={styles.content}>
        <div className={styles.leftBar}>
          <div className={styles.upload}>
            <Dropzone onUpdate={dropUpdate} />
          </div>
          <div className={styles.search}>
            <h5> Search </h5>
            <input type="search" />
            <button type="button">Search</button>
          </div>
          <div className={styles.tags}>
            <h5>Tags</h5>
            <ul className={styles.tagList}>
              <li>
                <Link to="/">+</Link>
                <Link to="/">-</Link>
                <Link to="/">tamed</Link>
              </li>
            </ul>
          </div>
        </div>
        <div className={styles.postList}>
          {!loaded ? <div>loading...</div> : null}
          <PostList posts={posts} />
        </div>
        <div className={styles.postView}>
          <Routes>
            <Route path="/post/:postId" element={<Post />} />
          </Routes>
        </div>
      </div>
    </div>
  )
}
