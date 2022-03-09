<script lang="ts">
	import { onMount } from 'svelte'
	import { query, modify } from './api.ts'
	import Post from './Post.svelte'
	import Search from './Search.svelte'
	import TagList from './TagList.svelte'
	export let name: string

	let loaded = false
	let posts = []
	let incTags=[]
	let excTags=[]

	const updateList  = async() => {
		const result=await query({includeTags: incTags, excludeTags:excTags})
		posts=result.posts
		console.log(posts[0])
		loaded=true
	}

	const onAddTags = async (postId:string, includeTags:string[]) => {
		await modify(postId, includeTags,[])
		await updateList()
	}

	const onRmTag = async (postId:string, excludeTag:string) => {
		await modify(postId, [], [excludeTag])
		await updateList()
	}

	const onSearch = async (includeTags, excludeTags) => {
		incTags = includeTags
		excTags = excludeTags
		await  updateList()
	}

	onMount(updateList)
</script>

<main>

	<div class="container">
		<h1>Tamed - File server</h1>
		<div class="row">
			<div class="col-md-4">
				<div class="bd-example">
					<div class="mb-3">
						<label for="uploadInput" class="form-label"
							>Multiple file upload</label
						>
						<input
							type="file"
							class="form-control"
							id="uploadInput"
							multiple="multiple"
						/>
					</div>
					<div class="d-grid gap-2 d-md-flex justify-content-md-end">
						<button class="btn btn-primary me-md-2" id="uploadButton">
							Upload
						</button>
					</div>
				</div>
				<div class="bd-example">
					<Search bind:includeTags={incTags} bind:excludeTags={excTags} {onSearch}/>
				</div>
			</div>
			<div class="col-md-8">
				<div class="bd-example">
					{#if incTags.length === 0 && excTags.length  === 0}
						Recent
					{:else}
						<TagList includeTags={incTags} excludeTags={excTags}/>
					{/if}
				</div>
				<div>
						{#if loaded}
							{#each posts as post}
							<div class="bd-example">
							<Post link={"/media/"+post.link} id={post.id} date={post.ctime} tags={post.tags} onAddTags={onAddTags} onRmTag={onRmTag}/>
						</div>
							{/each}
						{/if}
					<div>
			</div>
		</div>
	</div>
</main>

<style>
	.bd-example {
		margin: 1rem 0 0;
		padding: 1.5rem;
		border: 1px solid #dee2e6;
		border-radius: 0.25rem;
	}

/*	main {
		/*text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
*/
</style>
