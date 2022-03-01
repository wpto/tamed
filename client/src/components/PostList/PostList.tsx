import React from 'react'
import styles from './PostList.module.css'
import { PostPreview } from '../../types/render'

interface Props {
  posts: PostPreview[]
}

export const PostList: React.FC<Props> = ({ posts }) => {
  return (
    <div className={styles.list}>
      <div className={styles.container}>
        {posts.map((p) => (
          <div className={styles.link}>
            <img className={styles.thumb} src={p.link} />
          </div>
        ))}
      </div>
    </div>
  )
}
