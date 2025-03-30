package store

func DefineName(n string) {
	name := n
	updateName(name)
}

func updateName(name string) {
	NestCds[0][2] = name
	NestCds[1][1] = name
}

var NestCds = [][]string{
	{"cds", "init", "nest-cds"},
	{"cd", "nest-cds"},
	{"npm", "i"},
	{"npm", "i", "typescript", "ts-node", "@types/node", "@nestjs/core", "@nestjs/common", "@nestjs/platform-express", "reflect-metadata", "rxjs"},
	{"npm", "i", "ts-node-dev", "--save-dev"},
	{"npm", "i", "-D", "tsx"},
	{"mkdir", "src"},
	{"cd", "src"},
	{"nest", "g", "module", "app", "--no-flat"},
}

var TsConf string = `{
  "compilerOptions": {
    "target": "ES2020",
    "module": "CommonJS",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "moduleResolution": "Node",
    "outDir": "dist",
    "rootDir": "src",
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true
  },
  "include": ["src/**/*.ts"],
  "exclude": ["node_modules"]
}`

var MainTs string = `import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import cds from '@sap/cds';

async function bootstrap() {
  console.log('Starting NestJS...');
  const app = await NestFactory.create(AppModule);

  const expressApp = app.getHttpAdapter().getInstance();

  cds.on('bootstrap', (capApp) => {
    console.log('CAP bootstrap started');
    expressApp.use(capApp);
  });

  await cds.serve('all').in(expressApp);

  await app.listen(4000);
  console.log('NestJS and CAP are running on http://localhost:4000');
}

bootstrap();`
