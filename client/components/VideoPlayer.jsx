import styles from '../styles/VideoPlayer.module.scss'
import { useEffect, useRef } from "react"
import HLS from 'hls.js'
import PropTypes from 'prop-types'
import { VideoOverlay } from './VideoOverlay'

export const VideoPlayer = ({ title, runtime, genre, description, directory, setShowPlayer }) => {
  const videoPlayerRef = useRef()
  
  useEffect(() => {
    const hls = new HLS({
      enableWorker: true,
      autoStartLoad: true,
    })

    hls.attachMedia(videoPlayerRef.current)
    hls.on(HLS.Events.MEDIA_ATTACHED, () => {
      hls.loadSource(`${process.env.NEXT_PUBLIC_BACKEND_HOST}/api/v1/explore/stream/${directory}`)
    })

    return () => {
      if (hls) {
        hls.destroy()
      }
    }
  }, [videoPlayerRef, directory])

  return (
    <div className={styles.container}>
      <VideoOverlay setShowPlayer={setShowPlayer} videoRef={videoPlayerRef} />
      <video className={styles.player} ref={videoPlayerRef} controls={true} autoPlay={true} />
    </div>
  )
}

VideoPlayer.propTypes = {
  title: PropTypes.string.isRequired,
  runtime: PropTypes.string.isRequired,
  genre: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  directory: PropTypes.string.isRequired,
  setShowPlayer: PropTypes.func.isRequired,
}
