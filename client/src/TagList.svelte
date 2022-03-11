<script lang="ts">
  export let includeTags : string[] = []
  export let excludeTags : string[] = []
  export let onClick : () => {}
</script>
  <div>
    {#each includeTags as tag}
      <span class="badge badge-success" on:click={onClick.bind(null, tag)}>{tag}</span>{' '}
    {/each}
    {#each excludeTags as tag}
      <span class="badge badge-danger" on:click={onClick.bind(null, "-"+tag)}>{tag}</span>{' '}
    {/each}
  </div>
