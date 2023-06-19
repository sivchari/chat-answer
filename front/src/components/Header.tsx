import React from 'react';

import { Header as MantineHeader, Title } from '@mantine/core';

import { createStyles } from '@mantine/core'

const useStyles = createStyles((theme, _params, getRef) => ({
    titleArea:{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
}))

const Header: React.FC = () => {
  const { classes } = useStyles();
  return(
    <MantineHeader m='sm' height={50}>
      <div className={classes.titleArea}>
        <Title size='h1'>
          Deep Dive into Go
        </Title>
      </div>
    </MantineHeader>
  )
};

export default Header;
