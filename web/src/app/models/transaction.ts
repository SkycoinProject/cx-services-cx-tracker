import {Server} from './server';

export class Transaction {
  chainType: string;
  config: any;
  updatedAt: string;
  createdAt: string;
  hash: string;
  id: number;
  servers: Server[];
}
