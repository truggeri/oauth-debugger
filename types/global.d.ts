// see https://github.com/Microsoft/TypeScript/issues/3180#issuecomment-283089923

import * as _config from "CONFIG";

declare global {
  const CONFIG: typeof _config;
}