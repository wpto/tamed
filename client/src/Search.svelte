<script lang="ts">
  import TagList from './TagList.svelte'
  import { searchField, searchTags, rmTag } from './search.ts'
  export let onSearch = () => {}

  const handleSearchButton = () => {
    onSearch($searchTags.include, $searchTags.exclude)
  }

  const handleRemoveTag = (tag) => {
    const field = $searchField
    $searchField = rmTag(field, tag)
  }
</script>

<div>
  <div class="mb-3">
    <label for="searchInput" class="form-label">Search</label
    ><input
      type="search"
      class="form-control"
      id="searchInput"
      bind:value="{$searchField}"
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
  <div>
    <TagList
      includeTags="{$searchTags.include}"
      excludeTags="{$searchTags.exclude}"
      onClick="{handleRemoveTag}"
    />
  </div>
</div>
