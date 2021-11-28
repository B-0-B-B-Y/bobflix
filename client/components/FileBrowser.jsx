import styles from '../styles/FileBrowser.module.scss'
import PropTypes from 'prop-types'
import { FileCard } from './FileCard'
import { TextField } from '@material-ui/core'
import { VideoPlayer } from './VideoPlayer'

export const FileBrowser = ({ fileList, error }) => {
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
        {/* {fileList && !error && (
          fileList.map((file) => {
            return (
              <FileCard {...file} key={file.title} />
            )
          })
        )} */}
        <VideoPlayer />
      </div>
    </div>
  )
}

FileBrowser.propTypes = {
  fileList: PropTypes.shape({
    title: PropTypes.string.isRequired,
    runtime: PropTypes.string.isRequired,
    genre: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
  }),
  error: PropTypes.shape({
    Error: PropTypes.string.isRequired
  })
}