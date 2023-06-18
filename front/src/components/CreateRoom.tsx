import React, { useState } from 'react';

import { Link, useNavigate } from 'react-router-dom';
import { Button, Container, Group, Paper, Text, TextInput } from "@mantine/core";
import { notifications } from '@mantine/notifications';

import { useClient } from 'src/client/client.ts'
import { ChatService } from 'src/proto/proto/chat_connectweb.ts';


const CreateRoom: React.FC = () => {
  const [roomName, setRoomName] = useState<string>('');
  const navigate = useNavigate();
  const client = useClient(ChatService);
  const createRoom = async (): Promise<string> => {
    const res = await client.createRoom({ name: roomName });
    return res.id;
  };
  const delay = (ms: number) => new Promise(res => setTimeout(res, ms));
  const onClick = async (e: React.MouseEvent<HTMLButtonElement>) => {
    const id = await createRoom();
    setRoomName('');
    notifications.show({
      title: '作成完了',
      message: `${ roomName }が正常に作成されました😄`,
    })
    await delay(1500);
    navigate(`/room/${ id }`);
  }
  return(
    <Container size='xs'>
      <Paper shadow='md' m='xs' p='sm'>
        <TextInput mt='sm'
          label='ルーム名'
          withAsterisk
          placeholder='ルーム名を入力してください'
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setRoomName(e.target.value)}
        />
      </Paper>
      <Group position='right'>
        <Button mt='sm'>
          <Link to='/' style={{ textDecoration: 'none' }}>
            <Text color='white'>ルーム一覧に戻る</Text>
          </Link>
        </Button>
        <Button mt='sm' onClick={onClick}>送信</Button>
      </Group>
    </Container>
  )
}

export default CreateRoom;
