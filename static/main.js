;(function(){
  const renderTag = (tag, addclass="badge-secondary") => `<span class="badge ${addclass}">${tag}</span>`
  const renderFile = (post) => {

    const d = new Date(parseInt(post.ctime+'000'))

    const tags = post.tags.map(renderTag).join('')
    if (tags.length > 0) tags = `<div class="bd-example">${tags}</div>`
    return `<div class="bd-example">
    <a href="/media/${post.link}">${post.id}</a>
    ${d.toUTCString()}
    ${tags}
    <div><img style="max-width: 150px" src="/media/${post.link}"/></div>
  </div>`
  }

  const renderList = (list) => {
    console.log(list)
    const items = list.posts.map(renderFile).join('')
    return items
  }

  const query = async (incTags, excTags) => {
    const q = []
    if (incTags.length > 0 || excTags.length > 0)  {
      const val = [].concat(...incTags, ...(excTags.map(e => "-"+e))).join(' ')
      q.push("tags="+val)
    }

    let append = ''
    if (q.length > 0) {
      append = '?' + q.join('&')
    }

    const res = await fetch('/api/posts' + append)
    const json = await res.json()
    return json
  }

  const modifyAdd = async (id, incTag) => {
    const res = await fetch('/api/posts/'+id,  {
      method: 'PATCH',
      body: JSON.stringify({'add_tags':[incTag]})
    })
  }

  const renderSearchTags = (incTags, excTags) => {
    let tags = ''
    tags += incTags.map((tag) => renderTag(tag, "badge-success")).join(' ')
    if (excludeTags.length > 0) tags += ' '
    tags += excTags.map((tag) => renderTag(tag, "badge-danger")).join(' ')

    return tags
  }

  let handleAddTag =  () => {}

  const updatePostList = async (incTags, excTags) => {
    const list = await query(incTags, excTags)
    document.getElementById("searchTags").innerHTML = renderSearchTags(incTags, excTags)
    document.getElementById("searchFiles").innerHTML = renderList(list)
    const btns = document.querySelectorAll("button-add-tag")
    if (btns != null) {
        console.log('listen')
      for (let i = 0; i < btns.length; i++) {
        btns[i].addEventListener('click', handleAddTag)
      }
    }
  }

  let includeTags = []
  let excludeTags = []

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

  handleAddTag = async (e) => {
    const parent = e.target.parentNode.parentNode
    const input = parent.querySelector("input")
    const id = parent.getAttribute("data-id")
    if (input != null) {
      const value = input.value
      await modifyAdd(id, value)
      await updatePostList(includeTags, excludeTags)
    }
  }

  const handleSearchTags = (e) => {
    let query = e.target.value
    if (typeof query !== 'string') {
      query = ''
    }

    const tags = query.split(' ').filter(e => e.length !== 0 || e !== '-')
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

    document.getElementById("searchTagList").innerHTML = renderSearchTags(includeTags, excludeTags)
  }

  const handleSearchButton = async () => {
    updatePostList(includeTags, excludeTags)
  }

  const handleUpload = async (e) => {
    const input = document.getElementById("uploadInput")

    console.log(e)
    console.log(input.files)
    const files = input.files
    const data = new FormData()

    for (let i = 0; i < files.length; i++) {
      data.append('upload[]', files[i])
    }

    await fetch('/api/posts', {
      method: 'POST',
      body: data,
    })

    updatePostList(includeTags, excludeTags)
  }

  const start = async () => {
    updatePostList([], [])
    // document.getElementById("searchInput").addEventListener('input', handleSearchTags)
    // document.getElementById("searchButton").addEventListener("click", handleSearchButton)
    document.getElementById("uploadButton").addEventListener('click', handleUpload)
  }

  window.addEventListener('load', start)
})();