import {writable, derived} from 'svelte/store'

interface Search {
  field: string
  include: string[]
  exclude: string[]
}

export const searchField = writable<string>('')


export const searchTags = derived(searchField, ($searchField) => {
    const tags = $searchField
      .split(' ')
      .filter((e) => e.length > 0 && e !== '-')

    let includeTags = []
    let excludeTags = []
    for (let i = 0; i < tags.length; i++) {
      const t = tags[i]
      if (t[0] === '-') {
        excludeTags.push(t.slice(1))
      } else {
        includeTags.push(t)
      }
    }


    return {
      include: uniq(includeTags),
      exclude: uniq(excludeTags)
    }
})

  const uniq = (list: string[]) => {
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


