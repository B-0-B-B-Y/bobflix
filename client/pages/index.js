import { Loader } from '../components/Loader'
import { useState } from 'react'
import styles from '../styles/Home.module.scss'
import { FileBrowser } from '../components/FileBrowser'

export const Home = ({ videoList }) => {
  const [content, setContent] = useState(videoList)

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Bobflix</h1>
      <Loader show={content === null || content === undefined}/>
      <FileBrowser fileList={content} />
    </div>
  )
}

export const getServerSideProps = async (context) => {
  const response = await fetch(`${process.env.BACKEND_HOST || 'http://localhost:8080'}/api/v1/explore/list`, {
    method: 'GET',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    }
  })
  const videoList = await response.json()

  return {
    props: {
      videoList
    }
  }
}

export default Home
