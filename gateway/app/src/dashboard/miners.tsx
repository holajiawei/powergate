import React, {FunctionComponent, useState, useEffect} from 'react'
import clsx from 'clsx'
import { makeStyles } from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import Chart from './chart'
import Deposits from './Deposits'
import Orders from './Orders'
import MinersMeta, {Item} from './miners_meta'
import {APIClient} from '../_proto/miner_pb_service'
import {GetRequest} from '../_proto/miner_pb'

const useStyles = makeStyles(theme => ({
    paper: {
      padding: theme.spacing(2),
      display: 'flex',
      overflow: 'auto',
      flexDirection: 'column',
    },
    fixedHeight: {
      height: 240,
    },
  }))

const Miners: FunctionComponent = () => {
    const classes = useStyles()
    const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight)

    const [items, setItems] = useState<Item[]>([])
    useEffect(() => {
      const c = new APIClient("http://40.117.82.59:6002")
      const resp = c.get(new GetRequest(), (err, resp) => {
        if (err) {
          console.log("got error: ", err)
        }
        if (resp) {
          console.log("got miners response: ", resp.getIndex()?.toObject())
          const items = resp.toObject().index?.meta?.infoMap?.map(info => {
            const id = info[0]
            const meta = info[1]
            const item: Item = {
              id,
              online: meta.online,
              country: meta.location?.country || "unknown",
              latitude: meta.location?.latitude || -999,
              longitude: meta.location?.longitude || -999,
              userAgent: meta.useragent,
              lastUpdated: new Date(meta.lastupdated)
            }
            return item
          })
          setItems(items || [])
        }
      })
    }, [])

    return (
        <Grid container spacing={3}>
            {/* Chart */}
            {/* <Grid item xs={12} md={8} lg={9}>
              <Paper className={fixedHeightPaper}>
                <Chart />
              </Paper>
            </Grid> */}
            {/* Recent Deposits */}
            {/* <Grid item xs={12} md={4} lg={3}>
              <Paper className={fixedHeightPaper}>
                <Deposits />
              </Paper>
            </Grid> */}
            {/* Recent Orders */}
            <Grid item xs={12}>
              <Paper className={classes.paper}>
                <MinersMeta data={items} />
              </Paper>
            </Grid>
          </Grid>
    )
}

export default Miners