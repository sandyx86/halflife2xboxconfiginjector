
# Half-Life 2 Xbox Config Injector

## Usage:
1. Get the zip0_xbox.xzp file from your Xbox. Assuming you have loaded the game up at least once, it should be located at "Z:\hl2\hl2x".
2. Using the injection tool provided in this repository, inject your own config into the XZip file, this config will overwrite xboxuser.cfg.
3. Delete xboxuser.cfg, it should be located at "E:\TDATA\45410091\hl2\hl2x\cfg", if you do not do this the custom Xbox user config will not work.
4. Using the XZip file provided (should be called "modified.xzp") rename it to zip0_xbox.xzp and put it back into "Z:\hl2\hl2x". If it asks you to override the previous one, override it.

Sometimes the game will delete stuff from xboxuser.cfg so it's good to have a second config (for example: myconfig.cfg) and a bind in the config you injected into the xzp to exec the second config.
