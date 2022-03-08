;(function(){
  const renderTag = (tag, addclass="badge-secondary") => `<span class="badge ${addclass}">${tag}</span>`
  const renderFile = (post) => {
    const tags = post.tags.map(renderTag).join('')
    if (tags.length > 0) tags = `<div class="bd-example">${tags}</div>`
    return `<div class="bd-example">${tags}${post.id} ${post.ctime} <a href="/media/${post.link}">download</a></div>`
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

  const renderSearchTags = (incTags, excTags) => {
    let tags = ''
    tags += incTags.map((tag) => renderTag(tag, "badge-success")).join(' ')
    if (excludeTags.length > 0) tags += ' '
    tags += excTags.map((tag) => renderTag(tag, "badge-danger")).join(' ')

    return tags
  }

  const updatePostList = async (incTags, excTags) => {
    const list = await query(incTags, excTags)
    document.getElementById("searchTags").innerHTML = renderSearchTags(incTags, excTags)
    document.getElementById("searchFiles").innerHTML = renderList(list)
  }

  let includeTags = []
  let excludeTags = []

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

    document.getElementById("searchTagList").innerHTML = renderSearchTags(includeTags, excludeTags)
  }

  const handleSearchButton = async () => {
    updatePostList(includeTags, excludeTags)
  }

  const start = async () => {
    updatePostList([], [])
    document.getElementById("searchInput").addEventListener('input', handleSearchTags)
    document.getElementById("searchButton").addEventListener("click", handleSearchButton)
  }

  window.addEventListener('load', start)
})();