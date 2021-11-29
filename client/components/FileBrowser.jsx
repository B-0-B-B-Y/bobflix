import styles from '../styles/FileBrowser.module.scss'
import PropTypes from 'prop-types'
import { FileCard } from './FileCard'
import { TextField } from '@material-ui/core'
import { useState, useEffect } from 'react'
import { VideoPlayer } from './VideoPlayer'

export const FileBrowser = ({ videoList, error }) => {
  const [movieCards, setMovieCards] = useState(videoList)
  const [showPlayer, setShowPlayer] = useState(false)
  const [movieToPlay, setMovieToPlay] = useState(null)

  useEffect(() => {
    if (movieToPlay) {
      setShowPlayer(true)
    }

    return () => {
      setShowPlayer(false)
    }
  }, [movieToPlay])

  if (!showPlayer) {
    return (
      <div className={styles.container}>
        <TextField
          className={styles.search}
          id='standard-basic'
          label='Search for media...'
          inputProps={{
            style: { color: 'white' }
          }}
        />
        <div className={styles.content}>
          {movieCards && !error && (
            movieCards.map((video) => {
              return (
                <FileCard key={video.title} {...video} setMovieToPlay={setMovieToPlay} />
              )
            })
          )}
        </div>
      </div>
    )
  } else {
    return <VideoPlayer {...movieToPlay} setShowPlayer={setShowPlayer} />
  }
}

FileBrowser.propTypes = {
  videoList: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      title: PropTypes.string.isRequired,
      runtime: PropTypes.string.isRequired,
      genre: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
      directory: PropTypes.string.isRequired
    })
  ),
  error: PropTypes.shape({
    Error: PropTypes.string
  })
}