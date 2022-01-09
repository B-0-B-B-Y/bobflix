import { Loader } from '../components/Loader'
import styles from '../styles/Home.module.scss'
import { FileBrowser } from '../components/FileBrowser'

export const Home = ({ videoList, isEmpty, error }) => {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Bobflix</h1>
      <Loader show={videoList === null && error === null}/>
      <FileBrowser videoList={videoList} isEmpty={isEmpty} error={error} />
    </div>
  )
}

export const getServerSideProps = async (context) => {
  let props = {
    videoList: null,
    isEmpty: true,
    error: null
  }

  try {
    const response = await fetch(`${process.env.BACKEND_HOST || 'http://localhost:8080'}/api/v1/explore/list`, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      }
    })
    if (!response.ok) {
      props.error = await response.json()
    } else {
      props.videoList = await response.json().then((data) => data?.movies)
      props.isEmpty = false
    }
  } catch (err) {
    props.error = err.error
  }

  return {
    props
  }
}

export default Home
