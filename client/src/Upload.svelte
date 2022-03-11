<script lang="ts">
  import { upload } from './api.ts'
  export let onUpload = () => {}

  let uploadInput = null

  let uploadMsg = []
  let showUploadMsg = false
  let timerID = 0
  let timerDuration = 20000
  const setMessage = (msg: string[]) => {
    clearTimeout(timerID)
    uploadMsg = msg
    showUploadMsg = true
    timerID = setTimeout(() => {
      showUploadMsg = false
    }, timerDuration)
  }

  const handleUpload = async () => {
    if (
      uploadInput != null &&
      uploadInput.files != null &&
      uploadInput.files.length !== 0
    ) {
      const result = await upload(uploadInput.files)
      const count = uploadInput.files.length
      const errors = []
      let successCount = 0
      for (let i = 0; i < result.length; i++) {
        const item = result[i]
        if (item.error != null) {
          const match = item.error.match(/\(([^)]*)\).*bad upload type/)
          if (match != null) {
            errors.push(`Bad upload type: ${match[1]}`)
          } else {

          errors.push(item.error)
        }
        } else {
          successCount++
        }
      }

      setMessage([...errors,`Successfully uploaded ${successCount} ${successCount === 1 ? 'file' : 'files'}` ])
      onUpload()
    }
  }
</script>

<div>
  <div class="mb-3">
    <label for="uploadInput" class="form-label">Multiple file upload</label>
    <input
      type="file"
      class="form-control"
      name="upload[]"
      id="uploadInput"
      multiple="multiple"
      bind:this="{uploadInput}"
    />
  </div>
  <div class="d-grid gap-2 d-md-flex justify-content-md-end mb-3">
    <button
      class="btn btn-primary me-md-2"
      id="uploadButton"
      on:click="{handleUpload}"
    >
      Upload
    </button>
  </div>
  <div>
    {#if showUploadMsg} {#each uploadMsg as msg}<div class="msg">{msg}</div>
    {/each } {/if}
  </div>
</div>

<style>
  .msg:not(:last-child) {
    margin-bottom: 1rem;
  }
</style>
