package winrt

// common
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IClosable
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IAsyncOperation`1
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.AsyncOperationCompletedHandler`1
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.AsyncStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IAsyncInfo
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.HResult
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.DateTime
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.TimeSpan
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.Deferral
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.DeferralCompletedHandler
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IReference`1

// advertisement
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementWatcherStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementWatcher -method-filter add_Received -method-filter remove_Received -method-filter add_Stopped -method-filter remove_Stopped -method-filter Start -method-filter Stop -method-filter get_Status -method-filter get_AllowExtendedAdvertisements -method-filter put_AllowExtendedAdvertisements -method-filter get_ScanningMode -method-filter put_ScanningMode -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementReceivedEventArgs -method-filter get_RawSignalStrengthInDBm -method-filter get_BluetoothAddress -method-filter get_Advertisement -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementWatcherStoppedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEManufacturerData
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisement -method-filter get_LocalName -method-filter put_LocalName -method-filter get_ServiceUuids -method-filter get_ManufacturerData -method-filter get_DataSections -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementDataSection -method-filter get_DataType -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementPublisher -method-filter get_Advertisement -method-filter Start -method-filter Stop -method-filter get_Status -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisementPublisherStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.Advertisement.BluetoothLEScanningMode

// bluetooth
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEDevice -method-filter FromBluetoothAddressAsync -method-filter FromBluetoothAddressWithBluetoothAddressTypeAsync -method-filter Close -method-filter get_ConnectionStatus -method-filter add_ConnectionStatusChanged -method-filter remove_ConnectionStatusChanged -method-filter get_BluetoothDeviceId -method-filter GetGattServicesWithCacheModeAsync -method-filter GetGattServicesAsync -method-filter GetConnectionParameters -method-filter RequestPreferredConnectionParameters -method-filter add_ConnectionParametersChanged -method-filter remove_ConnectionParametersChanged -method-filter GetConnectionPhy -method-filter add_ConnectionPhyChanged -method-filter remove_ConnectionPhyChanged -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothConnectionStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothAddressType
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothDeviceId -method-filter !FromId
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothCacheMode
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothError
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEConnectionParameters
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEPreferredConnectionParameters
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEPreferredConnectionParametersRequest
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEPreferredConnectionParametersRequestStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEConnectionPhy
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.BluetoothLEConnectionPhyInfo

//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattSession -method-filter FromDeviceIdAsync -method-filter get_MaintainConnection -method-filter put_MaintainConnection -method-filter get_CanMaintainConnection -method-filter Close -method-filter get_MaxPduSize -method-filter add_MaxPduSizeChanged -method-filter remove_MaxPduSizeChanged -method-filter get_SessionStatus -method-filter add_SessionStatusChanged -method-filter remove_SessionStatusChanged -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattSessionStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattSessionStatusChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattDeviceServicesResult -method-filter !get_ProtocolError
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattCommunicationStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattDeviceService -method-filter get_Uuid -method-filter Close -method-filter GetCharacteristicsAsync -method-filter GetCharacteristicsWithCacheModeAsync -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristicsResult -method-filter !get_ProtocolError
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristic -method-filter get_Uuid -method-filter get_CharacteristicProperties -method-filter WriteValueWithOptionAsync -method-filter WriteValueAsync -method-filter ReadValueWithCacheModeAsync -method-filter ReadValueAsync -method-filter WriteClientCharacteristicConfigurationDescriptorAsync -method-filter add_ValueChanged -method-filter remove_ValueChanged -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristicProperties
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattWriteOption
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattReadResult -method-filter !get_ProtocolError
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattClientCharacteristicConfigurationDescriptorValue
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattValueChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattClientNotificationResult
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalCharacteristic
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalCharacteristicParameters
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalCharacteristicResult
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalDescriptorParameters
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattLocalService
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattProtectionLevel
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattReadRequest
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattReadRequestedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattRequestState
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProvider
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProviderAdvertisementStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProviderAdvertisingParameters
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattServiceProviderResult
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattSubscribedClient
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattWriteRequest
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Devices.Bluetooth.GenericAttributeProfile.GattWriteRequestedEventArgs

// event
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.TypedEventHandler`2
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.EventRegistrationToken

// buffer
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.IBuffer
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.Buffer -method-filter !CreateCopyFromMemoryBuffer -method-filter !CreateMemoryBufferOverIBuffer
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.IDataReader -method-filter ReadBytes -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.DataReader -method-filter FromBuffer -method-filter ReadBytes -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.IDataWriter -method-filter WriteBytes -method-filter DetachBuffer -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.DataWriter -method-filter WriteBytes -method-filter DetachBuffer -method-filter DataWriter -method-filter Close -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Storage.Streams.IRandomAccessStreamReference

// vector
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.Collections.IVector`1
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.Collections.IVectorView`1

// media
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.MediaPlaybackAutoRepeatMode

// media control
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSession
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionMediaProperties
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionPlaybackInfo
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionPlaybackControls
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionPlaybackStatus
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.GlobalSystemMediaTransportControlsSessionTimelineProperties
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.CurrentSessionChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.MediaPropertiesChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.PlaybackInfoChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.SessionsChangedEventArgs
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Media.Control.TimelinePropertiesChangedEventArgs

// credential locker
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Security.Credentials.PasswordVault -method-filter Add -method-filter Remove -method-filter Retrieve -method-filter FindAllByResource -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Security.Credentials.PasswordCredential -method-filter CreatePasswordCredential -method-filter get_Resource -method-filter get_UserName -method-filter get_Password -method-filter RetrievePassword -method-filter !*

// dispatcher queue, for getting back onto the UI thread
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.System.DispatcherQueueController -method-filter get_DispatcherQueue -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.System.DispatcherQueue -method-filter TryEnqueue -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.System.DispatcherQueueHandler

// xaml hosting
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Hosting.WindowsXamlManager
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Hosting.DesktopWindowXamlSource -method-filter put_Content -method-filter Close -method-filter !*

// xaml core
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.UIElement -method-filter add_LostFocus -method-filter remove_LostFocus -method-filter put_Opacity -method-filter put_Visibility -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.FrameworkElement -method-filter put_MaxWidth -method-filter put_MaxHeight -method-filter get_Resources -method-filter put_RequestedTheme -method-filter put_Margin -method-filter put_Width -method-filter put_Height -method-filter put_MinWidth -method-filter put_MinHeight -method-filter put_Name -method-filter put_HorizontalAlignment -method-filter put_VerticalAlignment -method-filter FindName -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Thickness
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.CornerRadius
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.ResourceDictionary -method-filter !*

// resource dictionaries are IMap<Object, Object>, and string keys are boxed
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.Collections.IMap`2 -method-filter Insert -method-filter Lookup -method-filter HasKey -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.PropertyValue -method-filter CreateString -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.RoutedEventHandler
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.RoutedEventArgs -method-filter !*

// xaml enums
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.HorizontalAlignment
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.VerticalAlignment
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Visibility
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.ElementTheme
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.FocusState
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.TextWrapping
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Orientation

// xaml layout
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Panel -method-filter get_Children -method-filter put_Background -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.UIElementCollection -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Grid -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.StackPanel -method-filter put_Orientation -method-filter put_Spacing -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Border -method-filter put_BorderBrush -method-filter put_BorderThickness -method-filter put_Child -method-filter put_Background -method-filter put_Padding -method-filter put_CornerRadius -method-filter !*

// xaml controls
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.TextBlock -method-filter put_TextWrapping -method-filter put_Text -method-filter put_FontSize -method-filter put_Foreground -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Control -method-filter get_BorderBrush -method-filter Focus -method-filter put_Padding -method-filter put_FontSize -method-filter put_Foreground -method-filter put_Background -method-filter put_CornerRadius -method-filter put_BorderThickness -method-filter put_BorderBrush -method-filter put_IsEnabled -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.ContentControl -method-filter put_Content -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Primitives.ButtonBase -method-filter add_Click -method-filter remove_Click -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.Button -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.TextBox -method-filter put_Text -method-filter get_Text -method-filter put_PlaceholderText -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.IconElement -method-filter put_Foreground -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.FontIcon -method-filter put_Glyph -method-filter put_FontSize -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Controls.PasswordBox -method-filter get_Password -method-filter put_Password -method-filter put_PlaceholderText -method-filter !*

// xaml brushes
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Media.Brush -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Xaml.Media.SolidColorBrush -method-filter put_Color -method-filter !*
//go:generate go run github.com/saltosystems/winrt-go/cmd/winrt-go-gen -debug -class Windows.UI.Color
