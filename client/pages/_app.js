import '../styles/globals.scss'
import { ThemeProvider } from '@material-ui/styles'
import { createTheme } from '@material-ui/core'

const theme = createTheme({
  palette: {
    primary: {
      main: '#db0000'
    },
    secondary: {
      main: '#564d4d'
    }
  }
})

function MyApp({ Component, pageProps }) {
  return (
    <ThemeProvider theme={theme}>
      <Component {...pageProps} />
    </ThemeProvider>
  )
}

export default MyApp
