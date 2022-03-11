interface queryOpts {
  includeTags: string[]
  excludeTags: string[]
  offset: number
  limit: number
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
  next: boolean
}

export const query = async ({
  includeTags,
  excludeTags,
  offset = 0,
  limit = 20,
}: queryOpts) => {
  const q: string[] = []
  if (includeTags.length > 0 || excludeTags.length > 0) {
    const val = []
      .concat(...includeTags, ...excludeTags.map((e) => '-' + e))
      .join(' ')
    q.push('tags=' + val)
  }

  q.push('offset=' + offset)
  q.push('limit=' + limit)

  let append = ''
  if (q.length > 0) {
    append = '?' + q.join('&')
  }

  const res = await fetch('/api/posts' + append)
  const json: queryResult = await res.json()
  return json
}

export const modify = async (
  postId: string,
  includeTags: string[],
  excludeTags: string[]
) => {
  const body = {}

  if (includeTags.length > 0) {
    body['add_tags'] = includeTags
  }

  if (excludeTags.length > 0) {
    body['rm_tags'] = excludeTags
  }

  const res = await fetch('/api/posts/' + postId, {
    method: 'PATCH',
    body: JSON.stringify(body),
  })
  const json: { ok: string } = await res.json()
  if (json.ok == 'changed') return true
  return false
}

export const remove = async (postId: string) => {
  const res = await fetch('/api/posts/' + postId, {
    method: 'DELETE',
  })
  const json = await res.json()
  console.log(json)
}

export const upload = async (files) => {
  const data = new FormData()
  const count = files.length

  for (let i = 0; i < files.length; i++) {
    const f = files[i]
    data.append('upload[]', f)
  }

  try {
    const res = await fetch('/api/posts', {
      method: 'POST',
      body: data,
    })
    const response = await res.json()
    return response
  } catch (e) {
    return [{ error: '' + e }]
  }
}
