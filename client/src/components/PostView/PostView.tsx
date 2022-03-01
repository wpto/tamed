import React from 'react'
import { PostViewProps } from '../../types/render'

export const PostView: React.FC<PostViewProps> = ({ fullUrl }) => {
  return (
    <div>
      <img src={fullUrl} />
    </div>
  )
}
