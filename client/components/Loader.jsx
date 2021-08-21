import styles from '../styles/Loader.module.scss'
import PropTypes from 'prop-types'

export const Loader = ({ show }) => {
  return show ? <div className={styles.loader} /> : null
}

Loader.propTypes = {
  show: PropTypes.bool
}

Loader.defaultProps = {
  show: false
}
