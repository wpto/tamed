<script lang="ts">
  export let onUpload = () => {}

  let uploadInput = null

  let uploadMsg = ''
  let showUploadMsg = false
  let timerID = 0
  let timerDuration = 3000
  const setMessage = (msg: string) => {
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
      console.log(uploadInput.files)
      const files = uploadInput.files
      const data = new FormData()
      const count = files.length

      for (let i = 0; i < files.length; i++) {
        const f = files[i]
        data.append('upload[]', f)
      }

      try {
        const res = await fetch('/api/posts', {
          method: 'POST',
          body: data,
        })
        const response = await res.json()
        setMessage(`Successfully uploaded ${count} ${count === 1 ? 'file':'files'}`)
        onUpload()
      } catch (e) {
        setMessage(`Error when uploading: ${e}`)
      }
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
    {#if showUploadMsg}
      {uploadMsg}
    {/if}
  </div>
</div>
