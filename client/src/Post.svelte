<script type="ts">
  export let link = ""
  export let id = ""
  export let date =""
  export let tags = ["there", "hello", "world"]
  export let onAddTags = () => {}
  export let onRmTag = () => {}

  let addTagField=''

  const handleAddTag = () => {
    if (addTagField.length > 0) {
      const tokens = addTagField.split(' ').filter(e => e.length !== 0 && e[0] !== '-')
      const map = {}
      for (let i = 0; i < tokens.length; i++) {
        map[tokens[i]] = true
      }

      onAddTags(id, Object.keys(map))
    }
  }

  const  handleRmTag = (tag) => {
    return () => {
      onRmTag(id, tag)
    }
  }

  let d = ""
  if (date.length > 0) {
    d = (new  Date(parseInt(date+'000'))).toUTCString()
  }
</script>

<div>
  <a class="mb-3" href="{link}">{id}</a>
  {d}
  <div class="mb-3">
    <img style="max-width: 150px" src="{link}" alt="post with id {id}" />
  </div>

  {#if tags.length > 0}
  <div class="mb-3">
    {#each tags as tag}
      <span class="badge badge-secondary" on:click={handleRmTag(tag)}>{tag}</span>{' '}
    {/each}
  </div>
  {/if}
  <div class="input-group mb-3">
    <div class="input-group-prepend">
      <button class="btn btn-sm btn-primary me-md-2" on:click={handleAddTag}>Add Tags</button>
    </div>
    <input type="text" class="form-control form-control-sm" bind:value={addTagField}/>
  </div>
</div>
