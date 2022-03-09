<script lang="ts">
  import TagList from './TagList.svelte'
  export let onSearch = () => {}

  let searchField = ''
  let includeTags: string[] = []
  let excludeTags: string[] = []

  const uniq = (list) => {
    const result = []
    const mmap = {}
    for (let i = 0; i < list.length; i++) {
      if (mmap[list[i]] == null) {
        result.push(list[i])
        mmap[list[i]] = true
      }
    }
    return result
  }

  const handleSearchField = () => {
    const tags = searchField
      .split(' ')
      .filter((e) => e.length !== 0 || e !== '-')
    includeTags = []
    excludeTags = []
    for (let i = 0; i < tags.length; i++) {
      const t = tags[i]
      if (t[0] === '-') {
        excludeTags.push(t.slice(1))
      } else {
        includeTags.push(t)
      }
    }

    includeTags = uniq(includeTags)
    excludeTags = uniq(excludeTags)
  }

  const handleSearchButton = () => {
    onSearch(includeTags, excludeTags)
  }
</script>

<div>
  <div class="mb-3">
    <label for="searchInput" class="form-label">Search</label
    ><input
      type="search"
      class="form-control"
      id="searchInput"
      bind:value="{searchField}"
      on:input="{handleSearchField}"
    />
  </div>
  <div class="d-grid gap-2 d-md-flex justify-content-md-between mb-3">
    <div class="me-md-2">tip: -tag for excluding</div>
    <button
      class="btn btn-primary me-md-2"
      id="uploadButton"
      on:click="{handleSearchButton}"
    >
      Search
    </button>
  </div>
  <div><TagList {includeTags} {excludeTags} /></div>
</div>
