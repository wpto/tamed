<script lang="ts">
  export let tags: string[] = []
  export let includeTag = () => {}
  export let excludeTag = () => {}

  const incTag = (tag: string) => () => includeTag(tag)
  const excTag = (tag: string) => () => excludeTag(tag)
</script>

<ul class="list-group list-group-flush">
  {#each tags as tag}
  <li class="list-group-item">
    <span class="badge badge-success tag-button" on:click={incTag(tag)}>+</span>
    <span class="badge badge-danger tag-button" on:click={excTag(tag)}>-</span> {tag}
  </li>
  {/each}
</ul>

<style>
  .tag-button {
    cursor: pointer;
  }
</style>