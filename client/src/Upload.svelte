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
      const errors = []
      for (let i = 0; i < result.length; i++) {
        const item = result[i]
        if (item.error != null) {
          errors.push(item.error)
        }
      }

      if (errors.length > 0) {
        setMessage(errors)
      } else {
        setMessage([
          `Successfully uploaded ${count} ${count === 1 ? 'file' : 'files'}`,
        ])
      }
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
  <div class="d-grid gap-2 d-md-flex justify-content-md-end">
    <button
      class="btn btn-primary me-md-2"
      id="uploadButton"
      on:click="{handleUpload}"
    >
      Upload
    </button>
  </div>
  <div>
    {#if showUploadMsg} {#each uploadMsg as msg} {msg}<br /><br/>
    {/each } {/if}
  </div>
</div>
