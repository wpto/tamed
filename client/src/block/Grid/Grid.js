import './Grid.css'
export const Grid = ({ artColumns = [] }) => {
  console.log(artColumns)
  return (
    <div className="grid">
      {artColumns.map((col) => (
        <div className="grid__column">
          <div className="grid__column-container">
            {col.map((art) => (
              <img className="image" src={art.link} alt={art.name} />
            ))}
          </div>
        </div>
      ))}
    </div>
  )
}
