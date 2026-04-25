; Weather App Installer Script for Inno Setup

[Setup]
; Basic information
AppId={{WeatherApp-ABCD-1234-EFGH-5678}}
AppName=Weather Info App
AppVersion=1.0
AppPublisher=Weather Team
AppPublisherURL=https://github.com/lissymay/infopogoda
DefaultDirName={pf}\WeatherApp
DefaultGroupName=Weather App
UninstallDisplayIcon={app}\weather-cli.exe
Compression=lzma2
SolidCompression=yes
WizardStyle=modern
OutputDir=installer
OutputBaseFilename=WeatherApp_Setup
PrivilegesRequired=admin
AllowNoIcons=yes

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"
Name: "russian"; MessagesFile: "compiler:Languages\Russian.isl"

[Tasks]
Name: "desktopicon"; Description: "Create a desktop icon"; GroupDescription: "Additional icons:"
Name: "quicklaunchicon"; Description: "Create a Quick Launch icon"; GroupDescription: "Additional icons:"; Flags: unchecked

[Files]
; Main application files
Source: "installer_content\weather-cli.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "installer_content\run_weather.bat"; DestDir: "{app}"; Flags: ignoreversion
Source: "installer_content\config.yaml"; DestDir: "{app}\config"; Flags: ignoreversion
Source: "installer_content\config-pogoda.yaml"; DestDir: "{app}\config"; Flags: ignoreversion

[Icons]
; Start menu icons
Name: "{group}\Weather App (CLI)"; Filename: "{app}\weather-cli.exe"
Name: "{group}\Weather App (BAT)"; Filename: "{app}\run_weather.bat"
Name: "{group}\Uninstall Weather App"; Filename: "{uninstallexe}"

; Desktop icon
Name: "{autodesktop}\Weather App"; Filename: "{app}\run_weather.bat"; Tasks: desktopicon

; Quick Launch icon
Name: "{userappdata}\Microsoft\Internet Explorer\Quick Launch\Weather App"; Filename: "{app}\run_weather.bat"; Tasks: quicklaunchicon

[Run]
; Run the BAT file after installation (with checkbox)
Filename: "{app}\run_weather.bat"; Description: "Launch Weather App"; Flags: postinstall nowait skipifsilent unchecked

[UninstallDelete]
; Delete config folder on uninstall
Type: filesandordirs; Name: "{app}\config"