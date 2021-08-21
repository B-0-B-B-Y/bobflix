import styles from '../styles/FileBrowser.module.scss'
import PropTypes from 'prop-types'
import { FileCard } from './FileCard'
import { TextField } from '@material-ui/core'

export const FileBrowser = ({ fileList }) => {
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
        {fileList.map((file) => {
          return (
            <FileCard {...file} key={file.title} />
          )
        })}
      </div>
    </div>
  )
}

FileBrowser.propTypes = {
  fileList: PropTypes.arrayOf(PropTypes.shape({
    title: PropTypes.string.isRequired,
    runtime: PropTypes.string.isRequired,
    genre: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
  })).isRequired
}