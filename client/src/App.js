import { useMemo, useState } from 'react'
import './App.css'

const mockImages = (() => {
  const result = []
  const add = (width, height, num) => {
    result.push({
      width,
      height,
      name: '/mock/' + ('000' + num).slice(-3) + '.jpg',
    })
  }
  add(620, 413, 1)
  add(400, 600, 2)
  add(800, 1103, 3)
  add(1462, 2000, 4)
  add(800, 1480, 5)
  add(800, 1119, 6)
  add(800, 1159, 7)
  add(800, 1280, 8)
  add(800, 1203, 9)
  return [...result, ...result, ...result]
})()

export const App = () => {
  const [columnCount, setColumnCount] = useState(3)

  const columns = useMemo(() => {
    const acc = Array(columnCount).fill(0)
    const result = Array(columnCount)
      .fill(null)
      .map((e) => [])
    for (let i = 0; i < mockImages.length; i++) {
      let min = 0
      for (let j = 1; j < columnCount; j++) {
        if (acc[j] < acc[min]) min = j
      }
      result[min].push(mockImages[i])
      acc[min] += mockImages[i].height
    }
    return result
  }, [columnCount])

  return (
    <div className="App">
      <div className="grid">
        {columns.map((col) => (
          <div className="grid__column">
            <div className="grid__column-container">
              {col.map((item) => (
                <img className="image" src={item.name} alt={item.name} />
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
