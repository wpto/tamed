<script lang="ts">
  export let includeTags : string[] = []
  export let excludeTags : string[] = []
</script>
  <div>
    {#each includeTags as tag}
      <span class="badge badge-success">{tag}</span>{' '}
    {/each}
    {#each excludeTags as tag}
      <span class="badge badge-danger">{tag}</span>{' '}
    {/each}
  </div>
