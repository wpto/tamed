import { useMemo, useState, useEffect } from 'react'
import { RoundButton } from './components/RoundButton/RoundButton'
import './App.css'
import { useScrollPosition } from '@n8tb1t/use-scroll-position'
import { Grid } from './block/Grid/Grid'
import { ProfileHeader } from './block/ProfileHeader/ProfileHeader'

// const mockImages = (() => {
//   const result = []
//   const add = (width, height, num) => {
//     result.push({
//       width,
//       height,
//       name: '/mock/' + ('000' + num).slice(-3) + '.jpg',
//     })
//   }
//   add(620, 413, 1)
//   add(400, 600, 2)
//   add(800, 1103, 3)
//   add(1462, 2000, 4)
//   add(800, 1480, 5)
//   add(800, 1119, 6)
//   add(800, 1159, 7)
//   add(800, 1280, 8)
//   add(800, 1203, 9)
//   return [...result, ...result, ...result]
// })()

const mockImages = (() => {
  const result = []
  for (let i = 0; i < 100; i++) {
    const width = Math.min(Math.floor(Math.random() * 20) + 1, 20) * 100
    const height = Math.min(Math.floor(Math.random() * 20) + 1, 20) * 100
    const color = (
      '000000' +
      Math.floor(((width * height) / 4000000) * 16777215).toString(16)
    ).slice(-6)
    result.push({
      width,
      height,
      link: `https://via.placeholder.com/${width}x${height}.jpg/${color}/${color}`,
      name: 'hi',
    })
  }
  return result
})()

export const App = () => {
  const [navHidden, setNavHidden] = useState(false)
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
      acc[min] += mockImages[i].height / mockImages[i].width
      console.log(min, acc)
    }
    return result
  }, [columnCount])

  useEffect(() => {
    console.log(window.scrollY)
  }, [])
  useScrollPosition(() => {}, [], null, true, 300)

  return (
    <div className="App">
      <div className={`bottom-nav ${navHidden ? 'bottom-nav--hidden' : ''}`}>
        <RoundButton
          onClick={() => {
            setColumnCount((columnCount % 5) + 1)
          }}
        >
          <ion-icon name="home" />
        </RoundButton>
        <RoundButton>
          <ion-icon name="ios-search" />
        </RoundButton>
        <RoundButton>
          <ion-icon name="ios-heart" />
        </RoundButton>
      </div>
      <ProfileHeader />
      <Grid artColumns={columns} />
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
