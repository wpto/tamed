import React, { useState, useEffect, useCallback } from 'react'
import styles from './Dropzone.module.css'

export const Dropzone: React.FC = () => {
  const [advancedUpload, setAdvancedUpload] = useState(false)
  const [isDragOver, setDragOver] = useState(false)
  useEffect(() => {
    const div = document.createElement('div')
    setAdvancedUpload(
      'draggable' in div ||
        ('ondragstart' in div &&
          'ondrop' in div &&
          'FormData' in window &&
          'FileReader' in window)
    )
  }, [])

  const prevent: React.DragEventHandler = useCallback((e) => {
    e.preventDefault()
    e.stopPropagation()
  }, [])

  const dragOver: React.DragEventHandler = useCallback((e) => {
    e.preventDefault()
    e.stopPropagation()
    setDragOver(true)
  }, [])
  const dragLeave: React.DragEventHandler = useCallback((e) => {
    e.preventDefault()
    e.stopPropagation()
    setDragOver(false)
  }, [])

  const drop: React.DragEventHandler = useCallback((e) => {
    e.preventDefault()
    e.stopPropagation()
    e.dataTransfer.files
    setDragOver(false)
  }, [])

  return (
    <form
      className={`${advancedUpload ? styles.boxDnd : styles.box} ${
        isDragOver ? styles.isDragOver : ''
      }`}
      method="post"
      action="/api/posts"
      encType="multipart/form-data"
      onDrag={prevent}
      onDragStart={prevent}
      onDragEnd={dragLeave}
      onDragOver={dragOver}
      onDragEnter={dragOver}
      onDragLeave={dragLeave}
      onDrop={drop}
    >
      <div className="box__input">
        <input
          className={styles.inputFile}
          type="file"
          name="files[]"
          id="file"
          data-multiple-caption="{count} files selected"
          multiple
        />
        <label htmlFor="file">
          <strong>Choose a file</strong>
          <span className={advancedUpload ? styles.showDnd : styles.hideDnd}>
            {' '}
            or drag it here
          </span>
          .
        </label>
      </div>
      <div className={styles.messageHidden}>Uploading…</div>
    </form>
  )
}
