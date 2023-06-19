import { useMemo } from "react";
import { ServiceType } from "@bufbuild/protobuf";
import {
  createConnectTransport,
} from "@bufbuild/connect-web";
import { createPromiseClient, PromiseClient } from "@bufbuild/connect";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  // generate stream http client
});

export function useClient<T extends ServiceType>(service: T): PromiseClient<T> {
  return useMemo(() => createPromiseClient(service, transport), [service]);
}
