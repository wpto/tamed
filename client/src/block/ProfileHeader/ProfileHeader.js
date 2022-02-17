import './ProfileHeader.css'

export const ProfileHeader = ({
  bgLink = '/dragon.png',
  avatarLink = '/dragon.png',
  userName = '@kifuku',
  userLabel = '利違 八无',
  userDesc = 'Some long long description',
}) => {
  return (
    <div className="profile-header">
      <div className="profile-header__bg-container">
        <img
          className="profile-header__bg"
          src={bgLink}
          alt="profile background"
        />
      </div>
      <div className="profile-header__user-card user-card">
        <div className="user-card__top">
          <div className="user-card__avatar-container">
            <img className="user-card__avatar" src={avatarLink} alt="avatar" />
          </div>
          <div className="user-card__name">{userName}</div>
        </div>
        <div className="user-card__desc user-desc">
          <div className="user-desc__text">{userLabel}</div>
          <div className="user-desc__text">{userDesc}</div>
        </div>
      </div>
    </div>
  )
}
