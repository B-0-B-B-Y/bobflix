import { makeStyles } from '@material-ui/core/styles'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import CardContent from '@material-ui/core/CardContent'
import Button from '@material-ui/core/Button'
import PlayCircleFilledIcon from '@material-ui/icons/PlayCircleFilled'
import Typography from '@material-ui/core/Typography'
import PropTypes from 'prop-types'

const useStyles = makeStyles({
  root: {
    maxWidth: 275,
    flex: '1 1 0',
    margin: 16,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
  center: {
    display: 'flex',
    flexFlow: 'row nowrap',
    justifyContent: 'center',
  },
}, { name: "Mui_FileCard" })

export const FileCard = ({ title, runtime, genre, description, directory, setMovieToPlay }) => {
  const classes = useStyles();

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography className={classes.title} color='textSecondary' gutterBottom>
          {genre}
        </Typography>
        <Typography variant='h5' component='h2'>
          {title}
        </Typography>
        <Typography className={classes.pos} color='textSecondary'>
          {description}
        </Typography>
        <Typography variant='caption' component='p'>
          Runtime: <span style={{color: '#db0000'}}>{runtime}</span>
        </Typography>
      </CardContent>
      <CardActions className={classes.center}>
        <Button className={classes.playButton} onClick={() => setMovieToPlay({ title, runtime, genre, description, directory })} variant='contained' size='medium' color='primary' endIcon={<PlayCircleFilledIcon />}>Watch Now</Button>
      </CardActions>
    </Card>
  )
}

FileCard.propTypes = {
  title: PropTypes.string.isRequired,
  runtime: PropTypes.string.isRequired,
  genre: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  directory: PropTypes.string.isRequired,
  setMovieToPlay: PropTypes.func.isRequired
}
