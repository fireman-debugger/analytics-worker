// parser.js

import { DataFrame } from './data-frame';
import { Config } from './config';

class Parser {
  async parse(file: string): Promise<DataFrame> {
    const configFile = file.replace('.csv', '.config');
    const config = Config.load(configFile);
    const df = await DataFrame.fromCSV(file, config);
    return df;
  }
}

const parser = new Parser();

export { parser };
export { Parser };