
interface queryOpts {
  includeTags: string[]
  excludeTags: string[]
}

interface postResult {
  id: string
  ctime: string
  tags: string[]
  link: string
}

interface queryResult {
  posts: postResult[]
  tags: string[]
}

export const query = async ({includeTags, excludeTags}:queryOpts ) => {
    const q : string[] = []
    if (includeTags.length > 0 || excludeTags.length > 0)  {
      const val = [].concat(...includeTags, ...(excludeTags.map(e => "-"+e))).join(' ')
      q.push("tags="+val)
    }

    let append = ''
    if (q.length > 0) {
      append = '?' + q.join('&')
    }

    const res = await fetch('/api/posts' + append)
    const json : queryResult = await res.json()
    return json
}

export const modify = async (postId:string, includeTags: string[], excludeTags: string[]) => {
  const body = {}

  if (includeTags.length > 0) {
    body['add_tags'] = includeTags
  }

  if (excludeTags.length > 0) {
    body['rm_tags'] = excludeTags
  }

  const res = await fetch('/api/posts/'+postId, {
    method: 'PATCH',
    body: JSON.stringify(body)
  })
  const json: {ok:string} = await res.json()
  if (json.ok == "changed") return true
  return false
}