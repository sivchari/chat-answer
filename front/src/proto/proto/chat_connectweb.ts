// @generated by protoc-gen-connect-web v0.10.0 with parameter "target=ts"
// @generated from file proto/chat.proto (package api, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ChatRequest, ChatResponse, CreateRoomRequest, CreateRoomResponse, GetRoomRequest, GetRoomResponse, ListMessageRequest, ListMessageResponse, ListRoomResponse } from "./chat_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.ChatService
 */
export const ChatService = {
  typeName: "api.ChatService",
  methods: {
    /**
     * @generated from rpc api.ChatService.CreateRoom
     */
    createRoom: {
      name: "CreateRoom",
      I: CreateRoomRequest,
      O: CreateRoomResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.ChatService.GetRoom
     */
    getRoom: {
      name: "GetRoom",
      I: GetRoomRequest,
      O: GetRoomResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.ChatService.ListRoom
     */
    listRoom: {
      name: "ListRoom",
      I: Empty,
      O: ListRoomResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.ChatService.ListMessage
     */
    listMessage: {
      name: "ListMessage",
      I: ListMessageRequest,
      O: ListMessageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.ChatService.Chat
     */
    chat: {
      name: "Chat",
      I: ChatRequest,
      O: ChatResponse,
      kind: MethodKind.BiDiStreaming,
    },
  }
} as const;

