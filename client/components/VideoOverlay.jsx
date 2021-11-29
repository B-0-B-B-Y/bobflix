import BackArrow from '@material-ui/icons/KeyboardBackspace';
import RewindButton from '@material-ui/icons/Replay10'
import { useState, useEffect } from 'react'
import styles from '../styles/VideoOverlay.module.scss'
import PropTypes from 'prop-types'

export const VideoOverlay = ({ videoRef, setShowPlayer }) => {
  const isVideoPlaying = (video) => !!(video.currentTime > 0 && !video.paused && !video.ended && video.readyState > 2);
  const [videoIsPlaying, setVideoIsPlaying] = useState(true)

  const handleRewindButtonClick = () => {
    if (videoRef.current.currentTime <= 10) {
      videoRef.current.currentTime = 0
    } else {
      videoRef.current.currentTime += -10
    }
  }

  useEffect(() => {
    if (isVideoPlaying(videoRef.current)) {
      setVideoIsPlaying(true)
    } else {
      setVideoIsPlaying(false)
    }
  }, [videoRef, videoIsPlaying])

  return (
    <div className={styles.controls}>
      <BackArrow onClick={() => setShowPlayer(false)} className={styles.button} color='primary' />
      <RewindButton className={styles.button} onClick={handleRewindButtonClick} color='primary' />
    </div>
  )
}

VideoOverlay.propTypes = {
  setShowPlayer: PropTypes.func.isRequired,
}
