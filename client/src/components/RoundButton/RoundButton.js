import './RoundButton.css'
export const RoundButton = ({ children, onClick }) => {
  return (
    <button type="button" className="round-button" onClick={onClick}>
      {children}
    </button>
  )
}
