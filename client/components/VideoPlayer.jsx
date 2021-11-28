import styles from '../styles/VideoPlayer.module.scss'
import { useEffect, useRef } from "react"
import HLS from 'hls.js'

export const VideoPlayer = ({}) => {
  const videoPlayerRef = useRef()
  
  useEffect(() => {
    const hls = new HLS({
      enableWorker: true,
      autoStartLoad: true,
      debug: true
    })

    hls.attachMedia(videoPlayerRef.current)
    hls.on(HLS.Events.MEDIA_ATTACHED, () => {
      // This needs refactoring to be dynamic, this will happen once I refactor the homepage UI to implement movie cards
      hls.loadSource(`${process.env.NEXT_PUBLIC_BACKEND_HOST}/api/v1/explore/stream/Aliens/playlist.m3u8`)
    })

    return () => {
      if (hls) {
        hls.destroy()
      }
    }
  }, [videoPlayerRef])

  return (
    <div className={styles.container}>
      <video className={styles.player} ref={videoPlayerRef} controls={true} />
    </div>
  )
}