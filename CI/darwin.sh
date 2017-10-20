#!/bin/bash
set -ev

#check env
df -h
diskutil list

ls $HOME/*
du -sh $HOME/*

if [ "$QT_HOMEBREW" == "true" ]
then
  #download and install qt with brew
  brew update
  brew install qt5
  brew outdated qt5 || brew upgrade qt5
  ln -s /usr/local/Cellar/qt/5.9.1 $HOME/Desktop/Qt5.9.1
else
  #download and install qt
  QT=qt-unified-mac-x64-online
  curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT.dmg https://download.qt.io/official_releases/online_installers/$QT.dmg
  hdiutil attach -noverify -noautofsck -quiet /tmp/$QT.dmg
  QT=qt-unified-mac-x64-3.0.0-online
  if [ "$IOS" == "true" ] || [ "$IOS_SIMULATOR" == "true" ]
  then
    /Volumes/$QT/$QT.app/Contents/MacOS/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs IOS=true
  else
    /Volumes/$QT/$QT.app/Contents/MacOS/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
  fi
  diskutil unmountDisk disk1
  rm -f /tmp/$QT.dmg
  ln -s $HOME/Qt $HOME/Desktop
fi

#prepare env
sudo chown $USER /usr/local/bin
sudo chown $USER $GOROOT/pkg | true

#check env
df -h
diskutil list

ls $HOME/*
du -sh $HOME/*

exit 0