<script lang="ts">
	import { onMount } from 'svelte'
	import { query, modify } from './api.ts'
	import {searchField, searchTags} from './search.ts'
	import Post from './Post.svelte'
	import Search from './Search.svelte'
	import TagList from './TagList.svelte'
	import Upload from './Upload.svelte'
	import TagPool from './TagPool.svelte'
	import InfiniteScroll from 'svelte-infinite-scroll'
	export let name: string

	let loaded = false
	let posts = []
	let hasNext = false
	let incTags = []
	let excTags = []
	let currPage = 0
	let currLimit = 20

	let tagPool = []

	const mergeTagPools = (first, another: string[]): string[] => {
		const map = {}
		for (let i = 0; i < first.length; i++) {
			map[first[i]] = true
		}
		for (let i = 0; i < another.length; i++) {
			map[another[i]] = true
		}
		const list = Object.keys(map)
		list.sort()
		return list
	}

	const updateList = async (next = false) => {
		const result = await query({
			includeTags: incTags,
			excludeTags: excTags,
			offset: currPage,
			limit: currLimit,
		})
		if (next) {
			posts = [...posts, ...result.posts]
			tagPool = mergeTagPools(tagPool, result.tags)
		} else {
			posts = result.posts
			tagPool = mergeTagPools([], result.tags)
		}
		hasNext = result.next
		loaded = true
	}

	const onAddTags = async (postId: string, includeTags: string[]) => {
		await modify(postId, includeTags, [])
		await updateList()
	}

	const onRmTag = async (postId: string, excludeTag: string) => {
		await modify(postId, [], [excludeTag])
		await updateList()
	}

	const onSearch = async (includeTags, excludeTags) => {
		incTags = includeTags
		excTags = excludeTags
		currPage = 0
		await updateList()
	}

	const handleScrollLoad = async () => {
		if (hasNext) {
			currPage++
			await updateList(true)
		}
	}

	const onUpload = async () => {
		if (incTags.length === 0 && excTags.length === 0 && currPage === 0) {
			await updateList(false)
		}
	}

	const onIncludeTag = async (tag: string) => {
		if ($searchTags.include.indexOf(tag) === -1) {
			$searchField = $searchField + " " +  tag
		}
	}

	const onExcludeTag = async (tag: string) => {
		if ($searchTags.exclude.indexOf(tag) === -1) {
			$searchField = $searchField + " -" +  tag
		}
	}

	onMount(updateList)
</script>

<main>
	<div class="container">
		<h1>Tamed - Image Fileserver</h1>
		<div><a href="//github.com/pgeowng/tamed">Github</a></div>
		<div class="row">
			<div class="col-md-4">
				<div class="bd-example">
					<Upload onUpload="{onUpload}" />
				</div>
				<div class="bd-example">
					<Search
						bind:includeTags="{incTags}"
						bind:excludeTags="{excTags}"
						{onSearch}
					/>
				</div>
				<div class="bd-example tag-pool">
					<TagPool tags="{tagPool}" includeTag={onIncludeTag} excludeTag={onExcludeTag}/>
				</div>
			</div>
			<div class="col-md-8">
				<div class="bd-example">
					Looking for: {#if incTags.length !== 0 || excTags.length !== 0}
					<div class="searchTagList">
						<TagList includeTags="{incTags}" excludeTags="{excTags}" />
					</div>
					{:else} All {/if}
				</div>
				<div class="post-list">
					{#if loaded}
						{#if posts.length === 0} 
							It seems like there is no matching images
						{:else }
						{#each posts as post}
						<div class="bd-example">
							<Post link={"/media/"+post.link} id={post.id} date={post.ctime}
							tags={post.tags} onAddTags={onAddTags} onRmTag={onRmTag}/>
						</div>
						{/each}
					{/if}
					{:else}
						Loading...
					{/if}
					<InfiniteScroll on:loadMore="{handleScrollLoad}" window="{true}" />
				</div>
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

	.tag-pool {
		padding:  0;
	}

	.post-list {
		margin-bottom: 20rem;
	}

	.searchTagList {
		display: inline-block;
	}
</style>
